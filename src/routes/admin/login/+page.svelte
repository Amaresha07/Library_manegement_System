<script lang="ts">
	import { tick } from 'svelte';
	import { goto } from '$app/navigation';

	let name = '';
	let password = '';
	let confirmPassword = '';
	let isLoading = false;
	let isSignup = false;
	let error = '';

	const handleSubmit = async () => {
		if (isSignup && password !== confirmPassword) {
			alert('Passwords do not match');
			return;
		}

		isLoading = true;
		error = '';
		await tick();

		try {
			const endpoint = isSignup ? 'http://localhost:9090/register' : 'http://localhost:9090/login';

			const response = await fetch(endpoint, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ name, password })
			});

			const data = await response.json();

			if (response.ok) {
				alert(data.message || `${isSignup ? 'Signup' : 'Login'} successful`);
				goto('/crudbook');
			} else {
				error = data.message || `${isSignup ? 'Signup' : 'Login'} failed`;
				alert(error);
			}
		} catch (err) {
			console.error(err);
			error = 'An error occurred while connecting to the server.';
			alert(error);
		} finally {
			isLoading = false;
			await tick();
		}
	};

	const toggleSignup = () => {
		isSignup = !isSignup;
		confirmPassword = '';
		error = '';
	};
</script>

<div class="bg-gradient">
	<div class="container">
		<h2 class="text-center text-3xl font-bold" style="color: #4c51bf;">
			{isSignup ? 'Create an Account' : 'Welcome Back'}
		</h2>

		<p class="text-center text-gray-600">
			{isSignup ? 'Sign up to manage your library' : 'Login to access your account'}
		</p>

		{#if error}
			<div class="error">{error}</div>
		{/if}

		<form on:submit|preventDefault={handleSubmit}>
			<div>
				<label for="name">Name</label>
				<input type="text" id="name" bind:value={name} required placeholder="Enter your name" />
			</div>

			<div>
				<label for="password">Password</label>
				<input
					type="password"
					id="password"
					bind:value={password}
					required
					placeholder="Enter your password"
				/>
			</div>

			{#if isSignup}
				<div>
					<label for="confirmPassword">Confirm Password</label>
					<input
						type="password"
						id="confirmPassword"
						bind:value={confirmPassword}
						required
						placeholder="Confirm your password"
					/>
				</div>
			{/if}

			<button type="submit" disabled={isLoading}>
				{#if isLoading}
					<span class="spinner"></span> Processing...
				{:else}
					{isSignup ? 'Sign Up' : 'Login'}
				{/if}
			</button>
		</form>

		<p class="mt-4 text-center">
			{isSignup ? 'Already have an account?' : "Don't have an account?"}
			<button on:click={toggleSignup} style="color: #4c51bf; font-weight: bold;">
				{isSignup ? 'Login here' : 'Sign up here'}
			</button>
		</p>
	</div>
</div>

<style>
	/* Background with gradient */
	.bg-gradient {
		background: linear-gradient(135deg, #667eea, #764ba2);
		min-height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	/* Container with shadow and smooth transition */
	.container {
		background: #fff;
		border-radius: 12px;
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
		padding: 40px;
		width: 100%;
		max-width: 450px;
		transition: transform 0.3s ease-in-out;
	}

	.container:hover {
		transform: translateY(-5px);
	}

	/* Form input styling */
	input {
		width: 100%;
		padding: 12px;
		margin: 10px 0;
		border: 1px solid #ccc;
		border-radius: 6px;
		box-sizing: border-box;
		transition: border-color 0.3s ease;
	}

	input:focus {
		border-color: #4c51bf;
		outline: none;
	}

	/* Button styling with transition and spinner */
	button {
		width: 100%;
		background: #4c51bf;
		color: white;
		border: none;
		padding: 12px;
		border-radius: 6px;
		font-size: 16px;
		cursor: pointer;
		transition: background 0.3s;
	}

	button:hover {
		background: #3c366b;
	}

	button:disabled {
		background: #b0b0b0;
		cursor: not-allowed;
	}

	.spinner {
		border: 4px solid rgba(255, 255, 255, 0.3);
		border-top: 4px solid #fff;
		border-radius: 50%;
		width: 24px;
		height: 24px;
		animation: spin 1s linear infinite;
		display: inline-block;
		vertical-align: middle;
		margin-right: 8px;
	}

	@keyframes spin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}

	/* Error message */
	.error {
		background: #ffcccc;
		color: #e53e3e;
		padding: 10px;
		border-radius: 6px;
		margin-bottom: 15px;
	}
</style>
