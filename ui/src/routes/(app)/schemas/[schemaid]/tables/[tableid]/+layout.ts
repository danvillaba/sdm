import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ params, parent }) => {
	let { session } = await parent();

	return {
		breadcrumbs: [
			session.database || '',
			params.schemaid,
			params.tableid
		]
	};
};
