import { useEffect, useRef } from "react";

export default function ScrollToTop() {
  const myRef = useRef<any>(null);

  useEffect(() => {
    const timeOutID = setTimeout(() => {
      if (myRef && myRef.current) {
        myRef.current.scrollIntoView();
      }
    }, 0);
    return () => clearTimeout(timeOutID);
  }, [myRef]);

  return <div ref={myRef}></div>;
}
