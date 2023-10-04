import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const ssr = false;

export const load: PageLoad = () => {
	const token = sessionStorage.getItem('token');
	if (token) throw redirect(307, '/');
};
