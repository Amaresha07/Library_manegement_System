import { render, fireEvent } from '@testing-library/svelte';
import { expect, vi } from 'vitest';
import BookList from '../BookList.svelte'; // Update the path if your component has a different file name
import { supabase } from '$lib/supabaseClient';
import { goto } from '$app/navigation';

// Mock supabase authentication
vi.mock('$lib/supabaseClient', () => ({
	supabase: {
		auth: {
			getSession: vi.fn(() => Promise.resolve({ data: { session: null } })) // Simulate logged out state
		}
	}
}));

// Mock navigation
vi.mock('$app/navigation', () => ({
	goto: vi.fn()
}));

describe('Borrow and Return Buttons', () => {
	it('redirects to /login when not logged in', async () => {
		const { getByText } = render(BookList);

		// Click Borrow button
		const borrowButton = getByText('Borrow Book');
		await fireEvent.click(borrowButton);

		// Expect redirection to login
		expect(goto).toHaveBeenCalledWith('/login');

		// Click Return button
		const returnButton = getByText('Return Book');
		await fireEvent.click(returnButton);

		// Expect redirection again
		expect(goto).toHaveBeenCalledWith('/login');
	});

	it('confirms action when logged in', async () => {
		vi.mocked(supabase.auth.getSession).mockResolvedValueOnce({
			data: { session: { user: { id: '123', email: 'test@example.com' } } }
		});

		const { getByText } = render(BookList);

		// Click Borrow button
		window.alert = vi.fn(); // Mock alert
		const borrowButton = getByText('Borrow Book');
		await fireEvent.click(borrowButton);

		expect(window.alert).toHaveBeenCalledWith('You have successfully borrowed the book!');

		// Click Return button
		const returnButton = getByText('Return Book');
		await fireEvent.click(returnButton);

		expect(window.alert).toHaveBeenCalledWith('You have successfully returned the book!');
	});
});
