import type { App, AppReviews } from '@/types'

const BASE_URL = 'http://localhost:8080'

async function handleResponse<T>(res: Response): Promise<T> {
  if (!res.ok) {
    const message = await res.text()
    throw new Error(message || 'Request failed')
  }
  try {
    return (await res.json()) as T
  } catch {
    // @ts-expect-error - allow void return when endpoint has no json
    return undefined
  }
}

export async function fetchApps(): Promise<App[]> {
  const res = await fetch(`${BASE_URL}/apps`)
  return handleResponse<App[]>(res)
}

export async function addNewApp(appId: string): Promise<void> {
  const res = await fetch(`${BASE_URL}/apps/new`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ appId }),
  })
  await handleResponse<void>(res)
}

export async function fetchReviews(
  appId: string,
  page: number = 1,
  limit: number = 20,
  since: string = new Date().toISOString(),
): Promise<AppReviews> {
  const url = new URL(`${BASE_URL}/apps/${appId}/reviews`)
  url.searchParams.set('page', String(page))
  url.searchParams.set('limit', String(limit))
  url.searchParams.set('since', since)
  console.log('sinceis', since, url.toString())
  const res = await fetch(url.toString())
  return handleResponse<AppReviews>(res)
}

export async function pollReviews(appId: string): Promise<void> {
  const url = new URL(`${BASE_URL}/apps/${appId}/poll`)
  const res = await fetch(url.toString())
  await handleResponse<void>(res)
}
