import type { LayoutLoad } from './$types';
import type { InfoResponse } from 'pb/sdm/session/v1/session_service';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';

export const ssr = false;

export const load: LayoutLoad = async () => {
	const token = sessionStorage.getItem('token');
	const api = axios.create({
		baseURL: 'http://localhost:8080/api',
		headers: { Authorization: token }
	});
	const res = await api
		.get<InfoResponse>('/session/info')
		.catch(() => {
			sessionStorage.removeItem('token');
			throw redirect(307, '/login');
		});

	return {
		session: res.data,
		api
	};
};
