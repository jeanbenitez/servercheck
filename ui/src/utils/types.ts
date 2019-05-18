// Return the type of each argument in Function F, as an array
// tslint:disable-next-line:ban-types
export type ArgumentTypes<F extends Function> = F extends (...args: infer A) => any ? A : never;
