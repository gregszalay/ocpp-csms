import { User } from "firebase/auth";
import { collection, Firestore, onSnapshot, query } from "firebase/firestore";

export function setUpSnapshotListener(
  setterCallback: (resultItems: any) => void,
  collectionName: string,
  query: any,
  db: Firestore | null,
  userInfo: User | null
) {
  if (db && userInfo) {
    const myQuery = query;
    // update station list if there is a change in Firestore
    const unsubscribe = onSnapshot(
      myQuery,
      (querySnapshot: any) => {
        const resultItems: any = [];
        querySnapshot.forEach((doc: any) => {
          resultItems.push(doc.data());
        });
        setterCallback(resultItems);
      },
      (error) => console.log("{error} " + JSON.stringify(error))
    );
    return () => {
      unsubscribe();
    };
  }
}
