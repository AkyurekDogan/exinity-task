output "alb_arn" {
  value = aws_lb.this.arn
  description = "The ARN of the ALB"
}

output "alb_dns_name" {
  value = aws_lb.this.dns_name
  description = "The DNS name of the ALB"
}

output "target_group_arn" {
  value = aws_lb_target_group.this.arn
  description = "The ARN of the target group"
}
