/**
 * Centralized runtime configuration.
 * All API routes and server-side code should import from here
 * instead of reading process.env inline.
 */
export const config = {
  apiUrl: process.env.API_URL || 'http://localhost:8090',
} as const;
