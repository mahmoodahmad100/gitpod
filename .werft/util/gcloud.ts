import { exec } from './shell';
import { sleep } from './util';
import { getGlobalWerftInstance } from './werft';
import { DNS, Record } from '@google-cloud/dns';
import { GCLOUD_SERVICE_ACCOUNT_PATH } from '../jobs/build/const';

export async function deleteExternalIp(phase: string, name: string, region = "europe-west1") {
    const werft = getGlobalWerftInstance()

    const ip = getExternalIp(name)
    werft.log(phase, `address describe returned: ${ip}`)
    if (ip.indexOf("ERROR:") != -1 || ip == "") {
        werft.log(phase, `no external static IP with matching name ${name} found`)
        return
    }

    werft.log(phase, `found external static IP with matching name ${name}, will delete it`)
    const cmd = `gcloud compute addresses delete ${name} --region ${region} --quiet`
    let attempt = 0;
    for (attempt = 0; attempt < 10; attempt++) {
        let result = exec(cmd);
        if (result.code === 0 && result.stdout.indexOf("Error") == -1) {
            werft.log(phase, `external ip with name ${name} and ip ${ip} deleted`);
            break;
        } else {
            werft.log(phase, `external ip with name ${name} and ip ${ip} could not be deleted, will reattempt`)
        }
        await sleep(5000)
    }
    if (attempt == 10) {
        werft.log(phase, `could not delete the external ip with name ${name} and ip ${ip}`)
    }
}

function getExternalIp(name: string, region = "europe-west1") {
    return exec(`gcloud compute addresses describe ${name} --region ${region}| grep 'address:' | cut -c 10-`, { silent: true }).trim();
}

export function createDNSRecord(domain: string, dnsZone: string, IP: string, slice: string) {
    const werft = getGlobalWerftInstance()
    werft.log(slice, `Creating DNS record. domain: ${domain} DNSZone: ${dnsZone} IPAddress: ${IP}`)

    const dnsClient = new DNS({
        projectId: 'gitpod-dev',
        keyFilename: GCLOUD_SERVICE_ACCOUNT_PATH
    })
    const zone = dnsClient.zone(dnsZone)
    const record = new Record(zone, 'a', {
        name: domain,
        ttl: 300,
        signatureRrdatas: [IP],
        data: IP
    })

    zone.addRecords(record).then(rsp => {
        werft.log(slice, `DNS Record created`)
    }, err => {
        throw new Error(err);
    }).catch(err => {
        throw new Error(err);
    })
}