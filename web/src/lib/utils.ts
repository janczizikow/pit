export const getInt = (str: string, fallback: number): number => {
	let res = parseInt(str);
	if (res <= 0 || isNaN(res) || !Number.isFinite(res)) {
		res = fallback;
	}
	return res;
};
