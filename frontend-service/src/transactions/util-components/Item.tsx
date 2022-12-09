import * as React from "react";
import { ReactElement, useState } from "react";

import Typography from "@mui/material/Typography";
import { CircularProgress, Grid, Stack } from "@mui/material";

import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";

import { Transaction } from "../typedefs/Transaction";
import { MeterValueType } from "ocpp-messages-ts/types/TransactionEventRequest";
import { Unsubscribe } from "@mui/icons-material";
import appTheme from "../../app/theme/AppTheme";


/***************************************************************************/

interface Props {
  transaction: Transaction;
}

/***************************************************************************/

export default function TransactionItem(props: Props) {
  let meterValues = props.transaction.meterValues;
  let meterValuesLen = meterValues.length;
  let meterValuesLast: MeterValueType = meterValues[meterValuesLen - 1];

  return (
    <Card
      sx={{
        minWidth: 275,
        background: props.transaction.energyTransferInProgress
          ? appTheme.palette.secondary.dark
          : "white",
      }}
    >
      <CardContent>
        <Typography
          variant="h1"
          component="div"
          padding={1}
          align="left"
          sx={{
            fontSize: 20,
          }}
        >
          <Stack spacing={1} direction="column">
            <Stack spacing={1} direction="row">
              <Typography>{"Charging Station ID: "}</Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {props.transaction.stationId}
              </Typography>
            </Stack>
            <Stack spacing={1} direction="row">
              <Typography>{"Energy Transfer In Progress: "}</Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {props.transaction.energyTransferInProgress ? "TRUE" : "FALSE"}
              </Typography>
            </Stack>

            <Stack spacing={1} direction="row">
              <Typography>{"Energy Transfer Started: "}</Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {props.transaction.energyTransferStarted}
              </Typography>
            </Stack>
            <Stack spacing={1} direction="row">
              <Typography>{"Energy Transfer Stopped: "}</Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {props.transaction.energyTransferStopped}
              </Typography>
            </Stack>
            <Stack spacing={1} direction="row">
              <Typography>
                {meterValuesLast.sampledValue[0].measurand + ": "}
              </Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {meterValuesLast.sampledValue[0].value} <span> </span>
                {meterValuesLast.sampledValue[0].unitOfMeasure?.unit}
              </Typography>
            </Stack>
            <Stack spacing={1} direction="row">
              <Typography>
                {meterValuesLast.sampledValue[1].measurand + ": "}
              </Typography>
              <Typography
                sx={{
                  fontWeight: "bold",
                }}
              >
                {" "}
                {meterValuesLast.sampledValue[1].value} <span> </span>
                {meterValuesLast.sampledValue[1].unitOfMeasure?.unit}
              </Typography>
            </Stack>
          </Stack>
        </Typography>
      </CardContent>
    </Card>
  );
}
