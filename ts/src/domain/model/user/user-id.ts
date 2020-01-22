export class UserId {
  constructor(public readonly value: string) {
    if ([value.length < 4, value.length > 32].some(v => v)) {
      throw new Error('assertion error');
    }
  }
}
