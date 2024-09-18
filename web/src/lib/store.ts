import type { SubmissionsResponse } from './types';
import { derived, writable } from 'svelte/store';

export const PAGE_SIZE = 50;

export function createSubmissionsStore() {
	const isFetching = writable(false);
	const query = writable({ page: 1, class: '' });
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
		selected: derived(query, (q) => q.class),
		subscribe: data.subscribe,
		data: derived(data, (d) => d),
		listSubmissions,
		changePage: (p: number) => {
			if (p > 0) {
				query.update((q) => ({ ...q, page: p }));
			}
		},
		selectClass: (cls: string) => {
			query.update((q) => {
				if (q.class === cls) {
					return {
						page: 1,
						class: ''
					};
				} else {
					return { page: 1, class: cls };
				}
			});
		}
	};
}
