import type { SubmissionsResponse } from './types';

export const PAGE_SIZE = 50;
export const listSubmissions = async (query: {
	season: string | number;
	page: string | number;
	class: string;
	mode: string;
}): Promise<SubmissionsResponse> => {
	const res = await fetch(
		`/api/v1/seasons/${query.season}/submissions?page=${query.page}&size=${PAGE_SIZE}&class=${query.class}&mode=${query.mode}&sort=-tier,duration`
	);
	const json = await res.json();
	if (res.status >= 300) {
		throw json;
	} else {
		return json;
	}
};
