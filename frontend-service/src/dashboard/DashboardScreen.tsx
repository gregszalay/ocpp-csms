import * as React from "react";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { makeStyles } from "@mui/styles";
import useMediaQuery from "@mui/material/useMediaQuery";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Navigator from "./util-components/Navigator";
import DnsRoundedIcon from "@mui/icons-material/DnsRounded";
//import theme from "./AppTheme";
//import { BrowserRouter, Switch, Route } from 'react-router-dom';
import PeopleIcon from "@mui/icons-material/People";
import PermMediaOutlinedIcon from "@mui/icons-material/PhotoSizeSelectActual";
import PublicIcon from "@mui/icons-material/Public";
import SettingsEthernetIcon from "@mui/icons-material/SettingsEthernet";
import SettingsInputComponentIcon from "@mui/icons-material/SettingsInputComponent";
import TimerIcon from "@mui/icons-material/Timer";
import SettingsIcon from "@mui/icons-material/Settings";
import PhonelinkSetupIcon from "@mui/icons-material/PhonelinkSetup";
import EvStationTwoToneIcon from "@mui/icons-material/EvStationTwoTone";
import StationList from "../stations/StationsScreen";
import Button from "@mui/material/Button";
import { Container, Paper } from "@mui/material";
import { AnalyticsTwoTone } from "@mui/icons-material";
import { Outlet } from "react-router-dom";
import Header from "./util-components/Header";
import theme from "../app/theme/AppTheme";

const drawerWidth = 120;

interface Props {
  selectedMenuItem_fromContext: MenuItem;
  handleMenuClick_fromContext: () => void;
  theme: any;
  menu: any[];
  isSmUp: boolean;
}

function DashboardScreen(props: Props) {
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const isSmUp = useMediaQuery(props.theme.breakpoints.up("lg"));
  //const isSmUp = useMediaQuery("(max-width:899px)");


  const [categories, setCategories] = React.useState(props.menu);

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  React.useEffect(() => {
    console.log("props.isSmUp: " + isSmUp);
    console.log("mobileOpen: " + mobileOpen);
  });
  //const classes = useStyles(appTheme);

  return (
    <Box
      sx={{
        display: "flex",
        minHeight: "100vh",
        overflow: "hidden",
        maxHeight: "100vh",
      }}
    >
      <Box
        component="nav"
        //sx={{ width: { md: drawerWidth }, flexShrink: { sm: 0 } }}
        sx={{
          width: { lg: drawerWidth },
          flexShrink: { lg: 0 },
          overflow: "hidden",
          boxShadow: "2px 5px 2px 1px rgba(0, 0, 0, .2)",
          zIndex:10
        }}
      >
        {isSmUp ? null : (
          <Navigator
            theme={props.theme}
            categories={categories}
            selectedMenuItem_passedDown={props.selectedMenuItem_fromContext}
            handleMenuClick_passedDown={props.handleMenuClick_fromContext}
            PaperProps={{ style: { width: drawerWidth } }}
            variant="temporary"
            open={mobileOpen}
            onClose={handleDrawerToggle}
          />
        )}
        <Navigator
          theme={props.theme}
          categories={categories}
          selectedMenuItem_passedDown={props.selectedMenuItem_fromContext}
          handleMenuClick_passedDown={props.handleMenuClick_fromContext}
          PaperProps={{ style: { width: drawerWidth } }}
          // open={true}
          //sx={{ display: { md: "block", sm: "none", xs: "none" } }}
          //sx={{ display: { sm: "block", xs: "block", md: "block", lg: "block" } }}
          sx={{ display: { lg: "block", md: "none",sm: "none", xs: "none" } }}
        />
      </Box>

      <Box
        sx={{
          flex: 1,
          display: "flex",
          flexDirection: "column",
          overflowY: "hidden",
          overflowX: "hidden",
          minHeight: "100vh",
          maxHeight: "100vh",
        }}
      >
        <Header
          theme={props.theme}
          onDrawerToggle={handleDrawerToggle}
          selectedMenuItem_passedDown={props.selectedMenuItem_fromContext}
        ></Header>
        <Box
          component="main"
          sx={{ flex: 1, py: 0, px: 0, mt: "0px", overflowY: "auto",
          background: theme.palette.secondary.light }}
        >
          <Outlet />
        </Box>
      </Box>
    </Box>
  );
}

export default DashboardScreen;
