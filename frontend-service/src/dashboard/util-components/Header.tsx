import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import HelpIcon from "@mui/icons-material/Help";
import IconButton from "@mui/material/IconButton";
import Link from "@mui/material/Link";
import MenuIcon from "@mui/icons-material/Menu";
import NotificationsIcon from "@mui/icons-material/Notifications";
import Toolbar from "@mui/material/Toolbar";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";
import { useTheme } from "@mui/system";
import { Drawer } from "@mui/material";
import { withFirebaseContext } from "../../app/contexts/Firebase";
import UserPanel from "./UserPanel";
import { withFirebaseDataContext } from "../../app/contexts/FirebaseData";

interface HeaderProps {
  onDrawerToggle: () => void;
  theme: any;
  selectedMenuItem_passedDown: MenuItem;
}

export default function Header(props: HeaderProps) {
  const theme = useTheme(props.theme);
  const appbarHeight = 80;

  const [userDrawerOpen, setuserDrawerOpen] = React.useState(false);

  const toggleDrawer = () => setuserDrawerOpen(!userDrawerOpen);

  return (
    <React.Fragment>
      <AppBar
        component="div"
        //color="primary"
        position="sticky"
        elevation={1}
        sx={{
          boxShadow: "0px 2px 2px 1px rgba(0, 0, 0, .2)",
          zIndex: 10,
          height: appbarHeight,
          justifyContent: "center",
          background: theme.palette.primary.dark,
        }}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid sx={{ display: { xl: "none", lg:"none", md: "block",sm: "block", xs: "block" } }} item>
              <IconButton
                color="inherit"
                aria-label="open drawer"
                onClick={props.onDrawerToggle}
                edge="start"
              >
                <MenuIcon />
              </IconButton>
            </Grid>
            <Grid item xs>
              <Typography
                sx={{ fontWeight: "bold", letterSpacing: "5px" }}
                color={theme.palette.primary.contrastText}
                variant="h5"
                component="h1"
              >
                {props.selectedMenuItem_passedDown.id}
              </Typography>
            </Grid>
            <Grid item>
              <Tooltip title="Alerts • No alerts">
                <IconButton color="inherit">
                  <NotificationsIcon />
                </IconButton>
              </Tooltip>
            </Grid>
            <Grid item>
              <IconButton
                color="inherit"
                sx={{ p: 0.5 }}
                onClick={toggleDrawer}
              >
                <Avatar src="/static/images/avatar/1.jpg" alt="My Avatar" />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
        <Drawer anchor={"right"} open={userDrawerOpen} onClose={toggleDrawer}>
          {withFirebaseDataContext(UserPanel, {
            selectedMenuItem_passedDown: props.selectedMenuItem_passedDown,
            closeHandler: toggleDrawer,
          })}
        </Drawer>
      </AppBar>
    </React.Fragment>
  );
}
