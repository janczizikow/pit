import { handleErrorWithSentry } from '@sentry/sveltekit';
import * as Sentry from '@sentry/sveltekit';

Sentry.init({
	dsn: 'https://1c0972ddd54a84431755269b86382d0c@o199983.ingest.us.sentry.io/4508015206924288',
	tracesSampleRate: 0.2,
	enabled: !import.meta.env.DEV,
	integrations: [Sentry.browserTracingIntegration()],
	tracePropagationTargets: [/^\/api\//]
});

// If you have a custom error handler, pass it to `handleErrorWithSentry`
export const handleError = handleErrorWithSentry();
