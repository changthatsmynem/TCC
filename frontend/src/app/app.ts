import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

interface Post {
  author: string;
  avatar: string;
  createdAt: string;
  imagePath: string;
}

interface CommentItem {
  id: number;
  author: string;
  avatar: string;
  message: string;
  createdAt: string;
}

interface FeedResponse {
  pageTitle: string;
  post: Post;
  comments: CommentItem[];
}

@Component({
  selector: 'app-root',
  imports: [CommonModule, FormsModule],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  private readonly http = inject(HttpClient);
  private readonly apiBaseUrl = 'http://localhost:8080/api';

  protected readonly loading = signal(true);
  protected readonly submitting = signal(false);
  protected readonly errorMessage = signal('');
  protected readonly pageTitle = signal('IT 08-1');
  protected readonly commentDraft = signal('');
  protected readonly post = signal<Post>({
    author: 'Change can',
    avatar: 'C',
    createdAt: '16 October 2021 16:00',
    imagePath: '/image1.png'
  });
  protected readonly comments = signal<CommentItem[]>([]);

  ngOnInit(): void {
    this.loadFeed();
  }

  protected submitComment(): void {
    const message = this.commentDraft().trim();
    if (!message || this.submitting()) {
      return;
    }

    this.submitting.set(true);
    this.errorMessage.set('');

    this.http.post<CommentItem>(`${this.apiBaseUrl}/comments`, { message }).subscribe({
      next: (comment) => {
        this.comments.set([...this.comments(), comment]);
        this.commentDraft.set('');
        this.submitting.set(false);
      },
      error: () => {
        this.errorMessage.set('Unable to save your comment right now.');
        this.submitting.set(false);
      }
    });
  }

  protected updateDraft(value: string): void {
    this.commentDraft.set(value);
  }

  private loadFeed(): void {
    this.loading.set(true);
    this.errorMessage.set('');

    this.http.get<FeedResponse>(`${this.apiBaseUrl}/feed`).subscribe({
      next: (response) => {
        this.pageTitle.set(response.pageTitle);
        this.post.set(response.post);
        this.comments.set([...response.comments]);
        this.loading.set(false);
      },
      error: () => {
        this.errorMessage.set('Unable to load the page content.');
        this.loading.set(false);
      }
    });
  }
}
