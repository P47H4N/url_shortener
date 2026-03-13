import React, { useState, type FormEvent, type ChangeEvent } from 'react';
import axios from 'axios';

interface ShortenResponse {
  status: string;
  message: string;
  data: {
    long_url: string;
    short_url: string;
  };
  error?: string;
}

interface ShortenRequest {
  long_url: string;
}

const base = 'http://localhost:8080/'

const Home: React.FC = () => {
  const [longUrl, setLongUrl] = useState<string>('');
  const [shortUrl, setShortUrl] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const handleShorten = async (e: FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setShortUrl('');

    try {
      const response = await axios.post<ShortenResponse>(
        `${base}create`,
        { long_url: longUrl } as ShortenRequest
      );

      if (response.data.status === "success") {
        setShortUrl(response.data.data.short_url);
      } else {
        setError(response.data.error || 'Something went wrong');
      }
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Unable to connect to API';
      setError(msg);
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ padding: '50px', textAlign: 'center', fontFamily: 'Arial' }}>
      <h1>URL Shortener</h1>
      <form onSubmit={handleShorten}>
        <input
          type="url"
          placeholder="Enter your long URL here..."
          value={longUrl}
          onChange={(e: ChangeEvent<HTMLInputElement>) => setLongUrl(e.target.value)}
          required
          style={{ padding: '10px', width: '350px', borderRadius: '5px', border: '1px solid #ccc' }}
        />
        <button 
          type="submit" 
          disabled={loading} 
          style={{ padding: '10px 20px', marginLeft: '10px', cursor: 'pointer' }}
        >
          {loading ? 'Processing...' : 'Shorten URL'}
        </button>
      </form>

      {error && <p style={{ color: 'red', marginTop: '15px' }}>{error}</p>}

      {shortUrl && (
        <div style={{ 
          marginTop: '30px', 
          padding: '15px', 
          background: '#f4f4f4', 
          display: 'inline-block',
          borderRadius: '8px' 
        }}>
          <p style={{ margin: '0 0 10px 0' }}>Success! Here is your link:</p>
          <a 
            href={shortUrl.startsWith('http') ? shortUrl : `${base}${shortUrl}`} 
            target="_blank" 
            rel="noreferrer" 
            style={{ fontWeight: 'bold', color: '#007bff' }}
          >
            {base}+{shortUrl}
          </a>
        </div>
      )}
    </div>
  );
};

export default Home;