import * as React from "react";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import useMediaQuery from "@mui/material/useMediaQuery";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Link from "@mui/material/Link";
import { Copyright as cpIcon } from "@mui/icons-material";

function Copyright() {
  return (
    <Typography variant="body2" color="secondary" align="center" mb="5px">
      {"Copyright © "}
      <Link color="inherit" href="https://muszertechnika.hu/">
        Műszertechnika Holding Zrt.
      </Link>{" "}
      {new Date().getFullYear()}.
    </Typography>
  );
}

export default Copyright;