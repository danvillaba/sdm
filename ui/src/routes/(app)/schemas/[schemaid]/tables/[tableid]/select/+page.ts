import type { QueryResponse } from 'pb/sdm/inspector/v1/inspector_service';
import type { PageLoad } from './$types';
import knex from 'knex';

export const load: PageLoad = async ({ params, parent }) => {
	const { api } = await parent();
	const schema = params['schemaid'];
	const table = params['tableid'];

	const b = knex({ client: 'pg' });
	const sql = b.withSchema(schema).from(table).limit(100).toString();

	const res = await api.post<Required<QueryResponse>>(
		'/inspector/query',
		{ sql }
	);

	return res.data;
};
