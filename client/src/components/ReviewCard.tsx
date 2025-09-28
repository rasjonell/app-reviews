import type { Review } from '../types'

function Stars({ rating }: { rating: number }) {
  return (
    <div className="flex gap-0.5">
      {Array.from({ length: 5 }).map((_, i) => (
        <svg
          key={i}
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill={i < rating ? 'currentColor' : 'none'}
          stroke="currentColor"
          className={
            'h-4 w-4 ' + (i < rating ? 'text-amber-400' : 'text-gray-300')
          }
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1"
            d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.802 2.036a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.802-2.036a1 1 0 00-1.175 0l-2.802 2.036c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.88 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
          />
        </svg>
      ))}
    </div>
  )
}

export default function ReviewCard({ review }: { review: Review }) {
  const date = new Date(review.timestamp)
  return (
    <div className="rounded-xl border border-black/10 bg-white shadow-sm p-5 flex flex-col gap-2">
      <div className="flex items-start justify-between gap-3">
        <div>
          <h3 className="text-base font-semibold text-gray-900">
            {review.title}
          </h3>
          <div className="text-sm text-gray-500">by {review.author}</div>
        </div>
        <Stars rating={review.rating} />
      </div>
      <div className="flex flex-col justify-between h-full">
        <p className="text-gray-800 leading-relaxed">{review.content}</p>
        <div className="text-xs text-gray-500 mt-1">
          {Number.isNaN(date.valueOf())
            ? review.timestamp
            : date.toLocaleString()}
        </div>
      </div>
    </div>
  )
}
