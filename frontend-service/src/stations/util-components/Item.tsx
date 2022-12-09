import * as React from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Firebase from "../../app/apis/firebase/Firebase";
import { ChargingStation } from "../typedefs/ChargingStation";
import { Stack } from "@mui/material";
/***************************************************************************/

interface Props {
  station: ChargingStation;
}

/***************************************************************************/

export default function StationItem(props: Props) {
  return (
    <Card sx={{ minWidth: 275 }}>
      <CardContent>
        <Stack spacing={1} direction="column">
          <Typography variant="h6" component="div">
            {props.station.id}
          </Typography>
          <Typography component="div">
            {" "}
            Model: <span> </span>
            {props.station.model}
          </Typography>
          <Typography component="div">
            Model: <span> </span>
            {props.station.serialNumber}
          </Typography>
          <Typography component="div">
            Model: <span> </span>
            {props.station.vendorName}
          </Typography>
          <Typography component="div">
            firmwareVersion: <span> </span>
            {props.station.firmwareVersion}
          </Typography>
          <Typography component="div">
            lastBoot: <span> </span>
            {props.station.lastBoot}
          </Typography>
        </Stack>
      </CardContent>
    </Card>
  );
}
