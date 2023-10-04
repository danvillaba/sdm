import type { GetSchemaTablesResponse } from 'pb/sdm/v1/service';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, parent }) => {
	let { api, session } = await parent();
	const res = await api.get<GetSchemaTablesResponse>(
		`/inspector/schemas/${params.schemaid}/tables`
	);
	const tables = res.data.tables || [];

	return {
		tables,
		breadcrumbs: [session.database || '', params.schemaid]
	};
};
