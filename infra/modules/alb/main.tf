resource "aws_lb" "this" {
  name               = "${var.alb_name}"
  internal           = false
  load_balancer_type = "application"
  security_groups   = [var.security_group_id]
  subnets            = var.subnet_ids
  enable_deletion_protection = false
  idle_timeout {
    seconds = 60
  }

  enable_cross_zone_load_balancing = true
}

resource "aws_lb_target_group" "this" {
  name     = "${var.alb_name}-target-group"
  port     = 80
  protocol = "HTTP"
  vpc_id   = var.vpc_id
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.this.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "fixed-response"
    fixed_response {
      status_code = 200
      content_type = "text/plain"
      message_body = "Welcome to ALB"
    }
  }
}

resource "aws_lb_listener_rule" "allow_grpc" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 100

  action {
    type = "forward"
    target_group_arn = aws_lb_target_group.this.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/grpc/*"]
  }
}
