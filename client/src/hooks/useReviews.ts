import { useInfiniteQuery } from '@tanstack/react-query'
import { fetchReviews } from '@/lib/api'

export function useReviewsInfiniteQuery(
  appId: string,
  limit = 20,
  since?: string,
) {
  return useInfiniteQuery({
    queryKey: ['reviews', appId, limit, { since }],
    initialPageParam: 1,
    queryFn: ({ pageParam }) => fetchReviews(appId, pageParam, limit, since),
    getNextPageParam: (lastPage, _pages, lastPageParam) => {
      const hasMore = lastPage.reviews.length === limit
      return hasMore ? (lastPageParam as number) + 1 : undefined
    },
    enabled: !!appId,
  })
}
