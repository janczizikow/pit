import type { SubmissionsResponse } from './types';
import { derived, writable } from 'svelte/store';

export const PAGE_SIZE = 50;

type SubmissionsQuery = { page: number; class: string };

export function createSubmissionsStore(initialQuery: SubmissionsQuery) {
	const isFetching = writable(false);
	const query = writable(initialQuery);
	const data = writable<SubmissionsResponse>({ data: [], metadata: {} });

	const listSubmissions = async ({
		page,
		classQuery
	}: {
		page: string | number;
		classQuery: string;
	}) => {
		try {
			isFetching.set(true);
			const res = await fetch(
				`/api/v1/submissions?page=${page}&size=${PAGE_SIZE}&class=${classQuery}&sort=-tier,duration`
			);
			const json = await res.json();
			data.set(json);
		} catch {
			// TODO:
		}
		isFetching.set(false);
	};

	return {
		query,
		isFetching: isFetching,
		subscribe: data.subscribe,
		data: derived(data, (d) => d),
		listSubmissions
	};
}
