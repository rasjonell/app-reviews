import { useQuery, useQueryClient } from '@tanstack/react-query'
import { pollReviews } from '@/lib/api'

export function useRefreshQuery(appId: string) {
  const qc = useQueryClient()

  const query = useQuery<void, Error>({
    queryKey: ['refresh', appId],
    queryFn: () => pollReviews(appId),
    enabled: false,
    retry: false,
  })

  async function refresh() {
    if (!appId) return
    await query.refetch()
    await qc.invalidateQueries({ queryKey: ['reviews', appId] })
  }

  return {
    ...query,
    refresh,
    isRefreshing: query.isFetching,
  }
}
