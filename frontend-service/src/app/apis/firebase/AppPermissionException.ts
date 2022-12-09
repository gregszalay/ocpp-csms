export default class AppPermissionException {
  readonly message: string;
  readonly name = "AppPermissionException";

  constructor(message: string) {
    this.message = message;
  }

  toString = () => {
    return `${this.name}: "${this.message}"`;
  };
}
