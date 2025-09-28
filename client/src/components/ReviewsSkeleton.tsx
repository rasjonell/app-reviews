export default function ReviewsSkeleton() {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {Array.from({ length: 6 }).map((_, i) => (
        <div
          key={i}
          className="animate-pulse h-28 rounded-xl border border-black/10 bg-white p-4"
        >
          <div className="h-5 bg-gray-200 rounded w-1/2 mb-2" />
          <div className="h-4 bg-gray-100 rounded w-full mb-1" />
          <div className="h-4 bg-gray-100 rounded w-3/4" />
        </div>
      ))}
    </div>
  )
}
