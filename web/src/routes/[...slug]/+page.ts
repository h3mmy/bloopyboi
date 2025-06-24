import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { api } from '$lib/api';

export const load: PageLoad = async ({ fetch }) => {
	const [appInfo] = await Promise.all([api(fetch, 'GET', 'info')]);
	if ((await appInfo.status) !== 200) {
		let data = await appInfo.json();
		throw error(appInfo.status, data.status);
	}
	return {
		appInfo: await appInfo.json()
	};
};
