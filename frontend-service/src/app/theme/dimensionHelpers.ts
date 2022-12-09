export const px = (value: number): string => {
  return value + "px";
};
export const vw = (value: number): string => {
  return value + "vw";
};
export const vh = (value: number): string => {
  return value + "vh";
};
export const deg = (value: number): string => {
  return value + "deg";
};

export const toHHMMSS = function (millis: string) {
  var sec_num = Math.floor(parseInt(millis, 10) / 1000); // don't forget the second param
  var hours: number | string = Math.floor(sec_num / 3600);
  var minutes: number | string = Math.floor((sec_num - hours * 3600) / 60);
  var seconds: number | string = sec_num - hours * 3600 - minutes * 60;

  if (hours < 10) {
    hours = "0" + hours;
  }
  if (minutes < 10) {
    minutes = "0" + minutes;
  }
  if (seconds < 10) {
    seconds = "0" + seconds;
  }
  return hours + ":" + minutes + ":" + seconds;
};
