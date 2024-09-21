import type { SubmissionsResponse } from './types';
import { derived, writable } from 'svelte/store';

export const PAGE_SIZE = 50;

type SubmissionsQuery = { page: number; class: string; mode: string };

export function createSubmissionsStore(initialQuery: SubmissionsQuery) {
	const isFetching = writable(false);
	const query = writable(initialQuery);
	const data = writable<SubmissionsResponse>({ data: [], metadata: {} });

	const listSubmissions = async ({
		page,
		classQuery,
		mode
	}: {
		page: string | number;
		classQuery: string;
		mode: string;
	}) => {
		try {
			isFetching.set(true);
			const res = await fetch(
				`/api/v1/submissions?page=${page}&size=${PAGE_SIZE}&class=${classQuery}&mode=${mode}&sort=-tier,duration`
			);
			const json = await res.json();
			if (res.status >= 300) {
				throw json;
			} else {
				data.set(json);
			}
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
