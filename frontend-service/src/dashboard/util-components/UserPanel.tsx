import * as React from "react";
import { Button, CircularProgress, Stack } from "@mui/material";
import { useNavigate } from "react-router-dom";
import UserInfo from "./UserInfo";
import { User } from "firebase/auth";
import Firebase from "../../app/apis/firebase/Firebase";
import theme from "../../app/theme/AppTheme";

import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";
import { Transaction } from "../../transactions/typedefs/Transaction";
import { ChargingStation } from "../../stations/typedefs/ChargingStation";

/***************************************************************************/

interface Props {
  firebase: Firebase;
  stations: ChargingStation[];
  transactions: Transaction[];
  chargetokens: IdTokenType[];
  userInfo: User | null;
  selectedMenuItem_passedDown: MenuItem;
  closeHandler: () => {};
}

/***************************************************************************/

export default function UserPanel(props: Props) {
  const navigate = useNavigate();

  return (
    <Stack
      direction="column"
      justifyContent="center"
      alignItems="center"
      sx={{
        height: "100vh",
        width: "300px",
        bgcolor: theme.palette.secondary.light,
      }}
      spacing={2}
    >
      {props.userInfo ? (
        <Stack
          direction="column"
          justifyContent="center"
          alignItems="center"
          sx={{ height: "100vh", width: "300px" }}
        >
          <UserInfo userInfo={props.userInfo} />

          <Button
            color="info"
            sx={{ mt: 5 }}
            onClick={() => {
              props.firebase.signOut();
              navigate("/");
            }}
          >
            Log out
          </Button>
        </Stack>
      ) : (
        <Stack
          sx={{ height: "100vh", width: "300px" }}
          direction="column"
          justifyContent="center"
          alignItems="center"
        >
          <CircularProgress color="info" />
        </Stack>
      )}
    </Stack>
  );
}
