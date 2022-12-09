import * as React from "react";
import Firebase from "../app/apis/firebase/Firebase";

import { CircularProgress, Container, List, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";

import TransactionItem from "./util-components/Item";
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

export default function TransactionsScreen(props: Props) {
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
          Transactions
        </Typography>
        {props.transactions && props.transactions.length > 0 ? (
          props.transactions.map((transaction: any) => (
            <TransactionItem transaction={transaction} />
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
