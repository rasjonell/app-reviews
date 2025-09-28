import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { addNewApp, fetchApps } from '../lib/api'
import type { App } from '../types'

export function useAppsQuery() {
  return useQuery<App[], Error>({
    queryKey: ['apps'],
    queryFn: fetchApps,
  })
}

export function useAddAppMutation() {
  const qc = useQueryClient()
  return useMutation<void, Error, { appId: string }>({
    mutationFn: async ({ appId }) => addNewApp(appId),
    onSuccess: async () => {
      await qc.invalidateQueries({ queryKey: ['apps'] })
    },
  })
}
