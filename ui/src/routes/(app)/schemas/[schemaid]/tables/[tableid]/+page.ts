import type { GetTableInfoResponse } from 'pb/sdm/v1/service';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, parent }) => {
	let { api } = await parent();
	const route = `/inspector/schemas/${params.schemaid}/tables/${params.tableid}`;
	const res = await api.get<Required<GetTableInfoResponse>>(route);

	return res.data;
};
