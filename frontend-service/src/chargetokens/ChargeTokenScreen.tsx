import * as React from "react";

import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import { Button, CircularProgress, Container, Divider } from "@mui/material";

import ChargeToken from "./util-components/Item";
import Firebase from "../app/apis/firebase/Firebase";
import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";
import { Transaction } from "../transactions/typedefs/Transaction";
import { ChargingStation } from "../stations/typedefs/ChargingStation";

/***************************************************************************/

interface Props {
  firebase: Firebase;
  stations: ChargingStation[];
  transactions: Transaction[];
  chargetokens: IdTokenType[];
}

/***************************************************************************/

export default function ChargeTokenScreen(props: Props) {
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
          RFID Tokens
        </Typography>
        {props.chargetokens && props.chargetokens.length > 0 ? (
          props.chargetokens.map((token: IdTokenType) => (
            <ChargeToken token={token} />
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
