import * as React from "react";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import Firebase from "../app/apis/firebase/Firebase";
import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";
import { Transaction } from "../transactions/typedefs/Transaction";
import StationItem from "./util-components/Item";
import { CircularProgress, Container, List } from "@mui/material";
import { ChargingStation } from "./typedefs/ChargingStation";

/***************************************************************************/

interface Props {
  firebase: Firebase;
  stations: ChargingStation[];
  transactions: Transaction[];
  chargetokens: IdTokenType[];
}

/***************************************************************************/

export default function StationsScreen(props: Props) {
  return (
    <Container sx={{ maxWidth: "100%", minWidth: "100%" }}>
      <Stack mt={3} spacing={2}>
        <Typography
          variant="h1"
          component="div"
          padding={2}
          align="left"
          sx={{
            fontSize: 40,
          }}
        >
          Charging Stations
        </Typography>
        {props.stations && props.stations.length > 0 ? (
          props.stations.map((station: ChargingStation) => (
            <StationItem station={station} />
          ))
        ) : (
          <Stack
            sx={{ height: "50vh" }}
            direction="column"
            justifyContent="center"
            alignItems="center"
          >
            <CircularProgress color="info" />
          </Stack>
        )}
      </Stack>
    </Container>
  );
}
