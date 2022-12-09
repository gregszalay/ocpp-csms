import * as React from "react";

import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";
import { Stack } from "@mui/material";

interface Props {
  token: IdTokenType;
}

/***************************************************************************/

export default function ChargeTokenItem(props: Props) {
  return (
    <Card sx={{ minWidth: 275 }}>
      <CardContent>
        <Stack spacing={1} direction="column">
          <Typography component="div">
            {" "}
            idToken: <span> </span>
            {props.token.idToken}
          </Typography>
          <Typography component="div">
            type: <span> </span>
            {props.token.type}
          </Typography>
        </Stack>
      </CardContent>
    </Card>
  );
}
