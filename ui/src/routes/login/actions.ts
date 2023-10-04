import { goto } from '$app/navigation';
import type { AxiosError } from 'axios';
import type { ConnectResponse } from 'pb/sdm/session/v1/session_service';
import api from 'provider/api';
import type { EventHandler } from 'svelte/elements.js';

export const handleSubmit: EventHandler<
	SubmitEvent,
	HTMLFormElement
> = async (e) => {
	const data = new FormData(e.currentTarget);

	return api
		.postForm<ConnectResponse>('/session/connect', data)
		.then((res) => {
			const token = res.data.token;
			if (!token) return;
			sessionStorage.setItem('token', token);
			goto('/');
		})
		.catch((err: AxiosError) => alert(err.response?.data));
};
