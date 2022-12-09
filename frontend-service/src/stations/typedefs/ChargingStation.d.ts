export type ChargingStation = {
  id: string;
  serialNumber: string;
  model: string;
  vendorName: string;
  firmwareVersion: string;
  modem: ChargingStationModem;
  location: ChargingStationLocation;
  lastBoot: string;
};
