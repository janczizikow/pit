/** [MDN Reference](https://developer.mozilla.org/docs/Web/API/URL/canParse_static) */
export default function canParseUrl(url: string | URL, base?: string): boolean {
	if (!('canParse' in URL)) {
		try {
			// @ts-expect-error https://github.com/microsoft/TypeScript/issues/55623
			return !!new URL(url, base);
		} catch {
			return false;
		}
	} else {
		return URL.canParse(url, base);
	}
}
