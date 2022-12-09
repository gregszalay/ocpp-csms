import { MeterValueType } from "ocpp-messages-ts/types/TransactionEventRequest";

export type Transaction = {
  stationId: string;
  energyTransferInProgress: bool;
  energyTransferStarted: string;
  energyTransferStopped: string;
  meterValues: MeterValueType[];
};
