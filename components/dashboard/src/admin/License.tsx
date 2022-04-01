/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { PageWithSubMenu } from "../components/PageWithSubMenu";
import { adminMenu } from "./admin-menu";

import { LicenseContext } from "../license-context";
import { useContext } from "react";
import { UserContext } from "../user-context";
import { Redirect } from "react-router-dom";
import CheckBox from "../components/CheckBox";

export default function License() {
    // @ts-ignore
    const { license, setLicense } = useContext(LicenseContext);
    const { user } = useContext(UserContext);

    if (!user || !user?.rolesOrPermissions?.includes("admin")) {
        return <Redirect to="/" />;
    }

    const featureList = license?.features;
    const enabledFeatures = license?.enabledFeatures;

    return (
        <div>
            <PageWithSubMenu subMenu={adminMenu} title="License" subtitle="License information of your account.">
                <p className="text-base text-gray-500 pb-4 max-w-2xl">
                    {license?.valid
                        ? "You have a valid license associated with this account. Following are the details:"
                        : "You do not have a valid license associated with this account."}
                </p>
                <div className="flex flex-col lg:flex-row">
                    <div className="lg:pl-14">
                        <div className="mt-4">
                            <h4>License Key</h4>
                            <input type="text" className="w-full" disabled={true} value={license?.key} />
                            <h4>Features Enabled</h4>
                            {featureList &&
                                featureList.forEach((feat, i) => {
                                    <CheckBox
                                        key={feat}
                                        title={capitalizeInitials(feat)}
                                        desc=""
                                        checked={enabledFeatures?.includes(feat) || false}
                                        disabled={true}
                                    />;
                                })}
                        </div>
                    </div>
                    <div className="lg:pl-14">
                        <div className="mt-4">
                            <h4>License Type</h4>
                            <input
                                type="text"
                                className="w-full"
                                disabled={true}
                                value={license?.type ? capitalizeInitials(license.type) : ""}
                            />
                        </div>
                        <div className="mt-4">
                            <h4>Number of seats</h4>
                            <input type="text" className="w-full" disabled={true} value={license?.seats} />
                        </div>
                        <div className="mt-4">
                            <h4>Available seats</h4>
                            <input type="text" className="w-full" disabled={true} value={license?.availableSeats} />
                        </div>
                    </div>
                </div>
            </PageWithSubMenu>
        </div>
    );
}

function capitalizeInitials(str: string): string {
    return str
        .split("-")
        .map((item) => {
            return item.charAt(0).toUpperCase() + item.slice(1);
        })
        .join(" ");
}
