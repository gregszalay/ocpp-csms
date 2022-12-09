import * as React from "react";

import PeopleIcon from "@mui/icons-material/People";
import SettingsIcon from "@mui/icons-material/Settings";
import EvStationTwoToneIcon from "@mui/icons-material/EvStationTwoTone";
import ElectricCarIcon from "@mui/icons-material/ElectricCar";
import TapAndPlayIcon from "@mui/icons-material/TapAndPlay";
import DashboardIcon from "@mui/icons-material/Dashboard";
import PieChartIcon from "@mui/icons-material/PieChart";
import MapIcon from "@mui/icons-material/Map";
import BackupTableIcon from "@mui/icons-material/BackupTable";
import CreditScoreIcon from "@mui/icons-material/CreditScore";
import GroupIcon from "@mui/icons-material/Group";
import PersonIcon from "@mui/icons-material/Person";

const dashboardMenu: Menu[] = [
  {
    headerId: "",
    children: [
      {
        id: "",
        label: "List of charging stations",
        icon: <EvStationTwoToneIcon />,
        route: "/dashboard/stations/list",
      },
    ],
  },
  {
    headerId: "",
    children: [
      {
        id: "",
        label: "List of charging stations",
        icon: <ElectricCarIcon />,
        route: "/dashboard/transactions",
      },
    ],
  },
  {
    headerId: "",
    children: [
      {
        id: "",
        label: "List of RFIDs",
        icon: <TapAndPlayIcon />,
        route: "/dashboard/chargetokens/list",
      },
    ],
  },
];

export default dashboardMenu;
