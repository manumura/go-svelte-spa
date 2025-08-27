import type { User } from '$lib/types';
import appConfig from '../../lib/config';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }): Promise<{ users: User[] }> => {
  const url = `${appConfig.baseUrl}/api/users`;
  const response = await fetch(url);

  if (!response.ok) {
    return { users: [] };
  }

  const users = await response.json();

  return {
    users
  };
};
