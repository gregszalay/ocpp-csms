import * as React from "react";
import { Box, CircularProgress, Stack, Typography } from "@mui/material";
import { User } from "firebase/auth";
import { Firestore } from "firebase/firestore";

interface Props {
  userInfo: User;
}

export default function UserInfo(props: Props) {
  {
    props.userInfo.getIdTokenResult().then((result) => {
      console.log("getIdTokenResult claims: " + JSON.stringify(result.claims));
    });
  }

  return (
    <React.Fragment>
      <Box sx={{ m: 3 }}>
        <Stack spacing={4}>
          {props.userInfo.displayName ? (
            <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
              {props.userInfo.displayName}
            </Typography>
          ) : null}

          {props.userInfo.email ? (
            <Stack spacing={1}>
              <Typography sx={{ color: "text.secondary" }}>
                Email address:
              </Typography>
              <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
                {props.userInfo.email}
              </Typography>
            </Stack>
          ) : null}

          {props.userInfo.phoneNumber ? (
            <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
              {props.userInfo.phoneNumber}
            </Typography>
          ) : null}

          {props.userInfo.photoURL ? (
            <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
              {props.userInfo.photoURL}
            </Typography>
          ) : null}

          {props.userInfo.providerId ? (
            <Stack spacing={1}>
              <Typography sx={{ color: "text.secondary" }}>
                Authentication provider:
              </Typography>
              <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
                {props.userInfo.providerId}
              </Typography>
            </Stack>
          ) : null}

          {props.userInfo.uid ? (
            <Stack spacing={1}>
              <Typography sx={{ color: "text.secondary" }}>UID:</Typography>
              <Typography sx={{ color: "text.secondary", fontWeight: "bold" }}>
                {props.userInfo.uid}
              </Typography>
            </Stack>
          ) : null}
        </Stack>
      </Box>
    </React.Fragment>
  );
}
