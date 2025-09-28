import { createFileRoute, useParams } from '@tanstack/react-router'
import { useMemo, useState } from 'react'
import { useReviewsInfiniteQuery } from '../hooks/useReviews'
import { useRefreshQuery } from '../hooks/useRefresh'
import ReviewCard from '../components/ReviewCard'
import { Link } from '@tanstack/react-router'
import ReviewsSkeleton from '@/components/ReviewsSkeleton'

export const Route = createFileRoute('/reviews/$appId')({
  component: ReviewsPage,
})

function formatYMD(date: Date) {
  const y = date.getFullYear()
  const m = String(date.getMonth() + 1).padStart(2, '0')
  const d = String(date.getDate()).padStart(2, '0')
  return `${y}-${m}-${d}`
}

function twoDaysAgoYMD() {
  const d = new Date()
  d.setDate(d.getDate() - 2)
  return formatYMD(d)
}

function ReviewsPage() {
  const { appId } = useParams({ from: '/reviews/$appId' })
  const [sinceDate, setSinceDate] = useState<string>(twoDaysAgoYMD())
  const sinceISO = useMemo(() => new Date(sinceDate).toISOString(), [sinceDate])
  const {
    data,
    isLoading,
    isError,
    error,
    refetch,
    hasNextPage,
    fetchNextPage,
    isFetchingNextPage,
  } = useReviewsInfiniteQuery(appId, 20, sinceISO)
  const { refresh: handleRefresh, isRefreshing } = useRefreshQuery(appId)

  return (
    <div className="min-h-screen bg-slate-50">
      <div className="max-w-6xl mx-auto px-4 py-8 space-y-4">
        <div className="flex items-center justify-between">
          <div>
            {data?.pages?.[0]?.app && (
              <h1 className="text-2xl font-bold text-gray-900">
                Reviews for {data.pages[0].app.name}
              </h1>
            )}
          </div>
          <div className="flex items-center gap-2">
            <label className="text-sm text-gray-700" htmlFor="sinceDate">
              Since
            </label>
            <input
              id="sinceDate"
              type="date"
              value={sinceDate}
              onChange={(e) => setSinceDate(e.target.value)}
              className="px-3 py-2 rounded-md border border-gray-300 text-gray-700 bg-white"
            />
            <button
              onClick={handleRefresh}
              disabled={isRefreshing}
              className="px-3 py-2 rounded-md border border-gray-300 text-gray-700 hover:bg-gray-50 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isRefreshing ? 'Refreshing…' : 'Refresh'}
            </button>
            <Link
              to="/"
              className="px-3 py-2 rounded-md bg-blue-600 text-white hover:bg-blue-700"
            >
              Back to Apps
            </Link>
          </div>
        </div>

        {isLoading && <ReviewsSkeleton />}

        {isError && (
          <div className="flex items-center justify-between gap-3 bg-rose-50 border border-rose-200 text-rose-800 p-3 rounded-md">
            <div>{error?.message}</div>
            <button
              onClick={() => refetch()}
              className="px-3 py-1 rounded border border-rose-300 text-rose-700 hover:bg-rose-100"
            >
              Retry
            </button>
          </div>
        )}

        {!!data && !isLoading && !isError && (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            {data.pages
              .flatMap((p) => p.reviews)
              .map((r) => (
                <ReviewCard key={r.id} review={r} />
              ))}
          </div>
        )}

        {!!data && !isLoading && !isError && (
          <div className="flex justify-center py-4">
            {hasNextPage ? (
              <button
                onClick={() => fetchNextPage()}
                disabled={isFetchingNextPage}
                className="cursor-pointer px-4 py-2 rounded-md border border-gray-300 text-gray-700 hover:bg-gray-50 disabled:opacity-50"
              >
                {isFetchingNextPage ? 'Loading…' : 'Load more'}
              </button>
            ) : (
              <div className="text-gray-500 text-sm text-center">
                <p>No more reviews</p>
                <p>consider changing the since date</p>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}
