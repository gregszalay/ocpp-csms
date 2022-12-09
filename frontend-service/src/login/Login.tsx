import * as React from "react";
import TextField from "@mui/material/TextField";
import { useNavigate } from "react-router-dom";
import { Button, Stack } from "@mui/material";
import { User, UserCredential } from "firebase/auth";

export default function Login(props: any) {
  let navigate = useNavigate();
  const [email, setEmail] = React.useState("");
  const [emailIsValid, setEmailIsValid] = React.useState(true);
  const [password, setPassword] = React.useState("");
  const [userSubmit, setUserSubmit] = React.useState(false);

  function validateEmail() {
    const re = /^[a-zA-Z0-9]+@[a-zA-Z0-9]+.+[A-Za-z]+$/;
    let valid = re.test(email);
    setEmailIsValid(valid);
  }

  function handleSubmit(event: any) {
    event.preventDefault();
    validateEmail();
    setUserSubmit(true);
    if (email && emailIsValid && password)
      props.firebase
        .signInWithEmailAndPassword(email, password)
        .then((authUser: UserCredential) => {
          console.log("authUser: " + { ...authUser });
          props.refresh(authUser);
          navigate("/dashboard/stations/list/");
        })
        .catch((error: any) => {
          console.log("{error} " + { error });
          navigate("/error");
        });
  }

  return (
    <React.Fragment>
      <Stack
        direction="row"
        justifyContent="center"
        alignItems="center"
        spacing={2}
        sx={{ m: 1, height: "100vh", width: "100%" }}
      >
        <Stack
          direction="column"
          justifyContent="center"
          alignItems="center"
          spacing={2}
          component="form"
          onSubmit={handleSubmit}
          sx={{}}
          noValidate
          autoComplete="off"
        >
          <TextField
            id="email"
            label="email"
            variant="outlined"
            error={userSubmit && (!emailIsValid || !email)}
            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
              setEmail(event.target.value);
              if (userSubmit) validateEmail();
            }}
          />
          <TextField
            id="password"
            label="jelszó"
            variant="outlined"
            type="password"
            error={userSubmit && !password}
            onChange={(event: React.ChangeEvent<HTMLInputElement>) =>
              setPassword(event.target.value)
            }
          />
          <Button variant="outlined" type="submit">
            Log In
          </Button>
        </Stack>
      </Stack>
    </React.Fragment>
  );
}
