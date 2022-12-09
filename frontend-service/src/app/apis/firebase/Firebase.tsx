import { FirebaseApp, initializeApp } from "firebase/app";
import * as firebaseAuthApi from "firebase/auth";
import React from "react";
import {
  Auth,
  AuthCredential,
  EmailAuthCredential,
  EmailAuthProvider,
  User,
  UserCredential,
} from "firebase/auth";

import dotenv from "dotenv";
dotenv.config();
console.log(`Your port is ${process.env.PORT}`);
console.log(`Your projectId is ${process.env.REACT_APP_PROJECT_ID}`);
console.log("Your projectId is" + process.env.REACT_APP_PROJECT_ID);

const firebaseConfig = {
  apiKey: process.env.REACT_APP_API_KEY,
  authDomain: process.env.REACT_APP_AUTH_DOMAIN,
  projectId: process.env.REACT_APP_PROJECT_ID!,
  storageBucket: process.env.REACT_APP_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_MESSAGING_SENDER_ID,
  appId: process.env.REACT_APP_APP_ID,
  measurementId: process.env.REACT_APP_APP_ID_MEASUREMENT_ID,
};

export default class Firebase {
  readonly auth: Auth;
  readonly firebaseApi: any;
  readonly app: FirebaseApp;
  currentUserCredential: EmailAuthCredential = new EmailAuthCredential();

  constructor() {
    console.log("firebaseConfig.apiKey " + firebaseConfig.apiKey);
    this.app = initializeApp(firebaseConfig);
    this.auth = firebaseAuthApi.getAuth();
  }

  // *** Auth API ***
  createUserWithEmailAndPassword = (
    email: string,
    password: string
  ): [Promise<UserCredential>, string] => {
    return [
      firebaseAuthApi.createUserWithEmailAndPassword(
        this.auth,
        email,
        password
      ),
      password,
    ];
  };

  signInWithCredential = (
    auth: Auth,
    credential: AuthCredential
  ): Promise<UserCredential> => {
    return firebaseAuthApi.signInWithCredential(auth, credential);
  };

  signInWithEmailAndPassword = (
    email: any,
    password: any,
    initialPassword: string
  ): Promise<UserCredential> | null => {
    if (password !== initialPassword) {
      const result = firebaseAuthApi.signInWithEmailAndPassword(
        this.auth,
        email,
        password
      );
      this.currentUserCredential = EmailAuthProvider.credential(
        email,
        password
      );
      return result;
    } else return null;
  };

  signOut = () => firebaseAuthApi.signOut(this.auth);

  resetPassword = (/*auth: Auth,*/ email: any) =>
    firebaseAuthApi.sendPasswordResetEmail(/*auth,*/ this.auth, email);

  updatePassword = (password: any) =>
    firebaseAuthApi.updatePassword(password.currentUser!, password);

  public get userInfo(): User | null {
    return this.auth.currentUser;
  }
}
