import { ArgumentTypes } from './types';

type FetchArguments = ArgumentTypes<typeof fetch>;

/**
 * @function fetchJSON
 * @description A HTTP JSON fetch functional utility
 */
export const fetchJSON = <T>(...args: FetchArguments): Promise<T> => {
  return fetch(...args)
    .then((response) => response.json());
};
