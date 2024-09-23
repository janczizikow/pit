import type { SubmissionsResponse } from './types';
import { derived, writable } from 'svelte/store';
import preloaded from '$lib/assets/preloaded.json';

export const PAGE_SIZE = 50;

type SubmissionsQuery = { page: number; class: string; mode: string; season: number };

export function createSubmissionsStore(initialQuery: SubmissionsQuery) {
	const isFetching = writable(false);
	const query = writable(initialQuery);
	const data = writable<SubmissionsResponse>(preloaded);

	const listSubmissions = async ({
		season,
		page,
		classQuery,
		mode
	}: {
		season: string | number;
		page: string | number;
		classQuery: string;
		mode: string;
	}) => {
		try {
			isFetching.set(true);
			const res = await fetch(
				`/api/v1/seasons/${season}/submissions?page=${page}&size=${PAGE_SIZE}&class=${classQuery}&mode=${mode}&sort=-tier,duration`
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
