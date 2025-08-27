import appConfig from '../../../lib/config';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
  const { id } = params;

  if (!id) {
    return { user: null };
  }

  const url = `${appConfig.baseUrl}/api/users/${id}`;
  const response = await fetch(url);

  if (!response.ok) {
    return { user: null };
  }

  const user = await response.json();
  return { user };
};
