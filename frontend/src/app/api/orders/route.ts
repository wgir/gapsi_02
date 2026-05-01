import { NextResponse } from 'next/server';
import { config } from '@/lib/config';
import { cookies } from 'next/headers';

export async function POST(request: Request) {
  const body = await request.json();
  const cookieStore = await cookies();
  const token = cookieStore.get('access_token')?.value;

  if (!token) {
    return NextResponse.json({ message: 'Unauthorized' }, { status: 401 });
  }

  try {
    const response = await fetch(`${config.apiUrl}/v1/orders`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    const data = await response.json();
    return NextResponse.json(data, { status: response.status });
  } catch (error) {
    console.error('Orders proxy error:', error);
    return NextResponse.json({ message: 'Internal server error' }, { status: 500 });
  }
}
