/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

export interface LicenseValidationResult {
    valid: boolean;
    msg?: string;
    issue?: LicenseIssue;
}

export type LicenseIssue = "seats-exhausted";

export interface LicenseInfo {
    key: string;
    seats: number;
    availableSeats?: number;
    valid: boolean;
    validUntil: string;
    plan?: string;
    features?: Map<string, boolean>;
    type?: string;
}

export interface GetLicenseInfoResult {
    isAdmin: boolean;
    licenseInfo: LicenseInfo;
}

export enum LicenseFeature {
    CreateSnapshot = "create-snapshot",
    // room for more
}

export interface LicenseService {
    validateLicense(): Promise<LicenseValidationResult>;
    getLicenseInfo(): Promise<GetLicenseInfoResult>;
    adminGetLicense(): Promise<LicenseInfo>;
    licenseIncludesFeature(feature: LicenseFeature): Promise<boolean>;
}
