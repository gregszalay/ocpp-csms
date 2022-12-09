import {
  addDoc,
  collection,
  deleteDoc,
  doc,
  Firestore,
  getDoc,
  getFirestore,
  setDoc,
  updateDoc,
} from "firebase/firestore";
import Firebase from "./Firebase";

import AppPermissionException from "./AppPermissionException";
import { ExistingFirestoreCollection } from "./MTFirestore";
import { ChargingStation } from "../../../stations/typedefs/ChargingStation";
import { Transaction } from "../../../transactions/typedefs/Transaction";
import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";

/**
 * CRUD functions
 */

export type FirestoreCRUDParamsType = {
  firebase: Firebase;
  db?: Firestore | null;
  record: ChargingStation | Transaction | IdTokenType;
  collectionName: string;
};

export async function createRecordInDB(
  firestoreRecordParams: FirestoreCRUDParamsType,
  id: string
): Promise<boolean> {
  const { firebase, db, record, collectionName } = firestoreRecordParams;
  try {
    const db = getFirestore(firebase.app);
    if (db) {
      console.log("db: ", { ...db });
      console.log("new record: ", { ...record });
      const response: any = await setDoc(doc(db, collectionName, id), record);
      console.log("response: ", { ...response });
    } else {
      throw new AppPermissionException("DB IS NULL");
    }
  } catch (err) {
    console.error(
      "Failed to create " + collectionName + " in DB, error: ",
      err
    );
    return false;
  }
  return true;
}

export async function createRecordInDBWithoutId(
  firestoreRecordParams: FirestoreCRUDParamsType
): Promise<boolean> {
  const { firebase, db, record, collectionName } = firestoreRecordParams;
  try {
    const db = getFirestore(firebase.app);
    if (db) {
      console.log("db: ", { ...db });
      console.log("new record: ", { ...record });
      const response: any = await addDoc(
        collection(db, collectionName),
        record
      );
      const newId = (await getDoc(response)).id;
      const response2: any = await setDoc(doc(db, collectionName, newId), {
        ...record,
        Id: newId,
      });
      console.log("response: ", { ...response });
    } else {
      throw new AppPermissionException("DB IS NULL");
    }
  } catch (err) {
    console.error(
      "Failed to create " + collectionName + " in DB, error: ",
      err
    );
    return false;
  }
  return true;
}

export async function updateRecordInDB(
  firestoreRecordParams: FirestoreCRUDParamsType,
  id: string
): Promise<boolean> {
  const { firebase, record, collectionName } = firestoreRecordParams;
  try {
    const db = getFirestore(firebase.app);
    const response: any = await setDoc(doc(db, collectionName, id), record);
  } catch (err) {
    console.error(
      "Failed to update " + collectionName + " in DB, error: ",
      err
    );
    return false;
  }
  return true;
}

export async function deleteRecordInDB(
  firestoreRecordParams: FirestoreCRUDParamsType,
  id: string
): Promise<boolean> {
  const { firebase, record, collectionName } = firestoreRecordParams;
  try {
    const db = getFirestore(firebase.app);
    const response: any = await deleteDoc(doc(db, collectionName, id));
  } catch (err) {
    console.error(
      "Failed to delete  " + collectionName + " in DB, error:" + err
    );
    return false;
  }
  return true;
}
