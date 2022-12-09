import { createTheme, ThemeProvider } from "@mui/material/styles";
import createPalette from "@mui/material/styles/createPalette";
import type {} from "@mui/x-data-grid/themeAugmentation";

const theme = createTheme({
  palette: {
    primary: {
      main: "#32681d",
      light: "#619648",
      dark: "#003d00",
      contrastText: "#edf2f2",
    },
    secondary: {
      main: "#c0ca33",
      light: "#ffffcf",
      dark: "#a69b97",
      contrastText: "#060614",
    },
    info: {
      main: "#2394ec",
      light: "#4ea7ec",
      dark: "#1b69a6",
    },
    success: {
      main: "#4cad50",
      light: "#7ac97e",
      dark: "#367539",
    },
    divider: "rgba(4,4,4,0.12)",
    error: {
      main: "#f54336",
      light: "#f96a60",
      dark: "#a22b23",
    },
    background: {
      default: "#ffff56",
    },
  },
  typography: {
    fontFamily: [
      "Archivo",
      "sans-serif",
      /*'Exo',
      'sans-serif',*/
    ].join(","),
    h5: {
      fontWeight: 300,
      fontSize: 26,
      letterSpacing: 0.5,
    },
    h6: {
      fontWeight: 200,
      fontSize: 20,
      letterSpacing: 0.4,
    },
  },
  shape: {
    borderRadius: 8,
  },
  mixins: {
    toolbar: {
      minHeight: 28,
    },
  },
  components: {
    // Use `MuiDataGrid` on both DataGrid and DataGridPro
    MuiDataGrid: {
      styleOverrides: {
        root: {
          backgroundColor: "#fafafa",
        },
      },
    },
  },
});

export default theme;
