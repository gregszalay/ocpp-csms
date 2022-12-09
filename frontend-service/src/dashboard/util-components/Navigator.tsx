import * as React from "react";
import Divider from "@mui/material/Divider";
import Drawer, { DrawerProps } from "@mui/material/Drawer";
import List from "@mui/material/List";
import Box from "@mui/material/Box";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import { AppBar, Button, Stack, Typography } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import Copyright from "./Copyright";
import theme from "../../app/theme/AppTheme";
import SvgIcon, { SvgIconProps } from "@mui/material/SvgIcon";

const content = {
  dashboardHeader1: "OCPP",
  dashboardHeader2: "Web",
};

const classes = {
  listHeaders: {
    py: 2,
    px: 3,
    color: "#0000000",
  },
  myDrawer: {
    height: "100vh",
    //background: theme.palette.primary.dark,
    //boxShadow: "16px",
  },
  myStack: { height: "100vh", background: theme.palette.secondary.main },
  item: {
    py: "0px",
    px: 4,
    /*color: '#60666b',*/
    "&:hover, &:focus": {
      bgcolor: "rgba(155, 155, 255, 0.08)",
    },
  },
  itemCategory: {
    boxShadow: "0 -1px 0 rgb(155,255,255,0.1) inset",
    py: 0,
    px: 0,
  },
  divider: { mt: 2 },
};

interface Props {
  PaperProps?: any;
  variant?: any;
  open?: boolean;
  onClose?: () => void;
  theme?: any;
  categories: Menu[];
  selectedMenuItem_passedDown: MenuItem;
  handleMenuClick_passedDown: (menuItem: MenuItem) => void;
  sx?: any;
}

export default function Navigator(props: Props) {
  const { ...other } = props;
  const myCategories: Menu[] = props.categories;
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <Drawer sx={{ ...classes.myDrawer }} variant="permanent" {...other}>
      <AppBar
        component="div"
        color="secondary"
        position="sticky"
        //elevation={-1}
        sx={{
          zIndex: -1,
          height: 80,
          justifyContent: "center",
          bgcolor: theme.palette.secondary.main,
        }}
      >
        <Stack
          spacing={0}
          alignItems="center"
          justifyContent="center"
          direction="column"
        >
          <Typography
            align="center"
            color="inherit"
            variant="h6"
            component="h1"
            sx={{ fontWeight: "bold" }}
          >
            {content.dashboardHeader1}
          </Typography>
          <Typography
            align="center"
            color="inherit"
            variant="h6"
            component="h1"
            sx={{ fontWeight: "bold" }}
          >
            {content.dashboardHeader2}
          </Typography>
        </Stack>
      </AppBar>
      <Stack
        sx={{ ...classes.myStack }}
        direction="column"
        justifyContent="space-between"
      >
        <List disablePadding>
          <ListItem
            sx={{ ...classes.item, ...classes.itemCategory }}
          ></ListItem>
          {myCategories.map(({ headerId, children }) => (
            <Box key={headerId}>
              <ListItem sx={{ ...classes.listHeaders }}>
                <ListItemText>{headerId}</ListItemText>
              </ListItem>
              {children.map(({ id: name, icon, label, route }) => (
                <ListItem disablePadding key={name} >
                  <ListItemButton
                    selected={props.selectedMenuItem_passedDown.id === name}
                    sx={{ ...classes.item }}
                    onClick={() => {
                      navigate(location.pathname.split("/")[0] + route);
                      props.handleMenuClick_passedDown({
                        id: name,
                        label,
                        icon,
                        route,
                      });
                    }}
                  >
                    <SvgIcon sx={{ fontSize: 60 }}>{icon}</SvgIcon>
                    <ListItemText>{name} </ListItemText>
                  </ListItemButton>
                </ListItem>
              ))}
              <Divider sx={{ ...classes.divider }} />
            </Box>
          ))}
        </List>

        <Copyright />
      </Stack>
    </Drawer>
  );
}
